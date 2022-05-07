# Bolus Wizard reverse engineer
User Input
---
- Blood Sugar (BG) 8.5mmol/l 
-- BG correction to target blood sugar 2.0units
- Grams of Carbs 10grams
-- carb correction to target blood sugar 1.3units)
- Bolus value = 3.3 units of insulin
## Required Values
### Carb Ratio
- grams per unit of insulin
-- increments of 30 minutes
-- ex: 11:30am to 4pm 1gram per 6.6 units of insulin
-- `[{"ratio": 1/6.5},{"start":11:30am},{"end":4pm}]`
### ISF - insulin sensitivity factor
- mmool per unit of insulin
-- How many units of insulin is required to reduce the bg by 1mmol (ex: 6.0bg to 5.0bg)
-- if it is 2:30pm isf is 1.4, and your BG is 7.0 and your target is 5.7, the diffence is 1.3
-- So it would take 1.4 units to correct an overage of 1mmol
- 30 minute increments
-- ex: 12am to 6am 1.2units of insulin
-- `[{"sensitivity":1/1.2},{"start":12am},{"end":6am}]`
### Active insuelin time (calulate bolus adjustment from previous meal)
-- 15 minute increments
-- ex: active insulin time
-- `{"activeInsulinTime":60min}`
### BG Target
- acceptal blood sugar value between certain times
- increments of 30 minutes
-- ex: 12am to 6am 5.4bg to 6.6bg
-- use "high" value as target when blood sugar is high
-- `[{"start":12am},{"end":6am},{"high":6.6},{"low":5,4}]`

## Basal setup
### required values
- insulin units per hour ( accuracy of two decimal points)
- 30min increments
-- accuracy 2 decimal points
-- `[{"start":12am,"end":2am, "unitsPerHour":1.50}]`

```typescript
function bolusWizard(bg: float,carbs: int): float{
      // Have to find current time to calculate the BG correction value...
      // Have to find current time to know what is the target blood sugar value
      // Current time is 3am, ISF insulin sensitivity is 1.2units of insulin
      // to reduce BG by 1 value 1.2 units of insulin is required between 12am and 3am
      const time = Date.now()
      const bgTarget = getBGTarget(time);

      // BG Correction
      const ISF = getInsulinSensitivityFactor(time);
      const bgCorrection = getBgCorrection(bg, bgTarget, ISF)

      // Carb correction
      const carbRatio = getCarbRatio(time);
      const carbCorrection = carbRatio * carbs;

      // Correct the correction base on boluses made in the past
      const activeInsulin = getActiveInsulin();

      const totalCorrection = carbCorrection + bgCorrection - activeInsulin;

      // round response to 2 decimals
      return totalCorrection;
};

function getInsulinSensitivityFactor(time): float{
      // look up what the ISF value should be based on the current time the function is being called
      // Lets say its 3am
      // maybe use an array of object, iterate over the array, look at each position , and if the current time
      // falls between the start and end field, return that value.

      const result = isfList.find(isf => {if(isf.start >= time && isf.end > time) return isf.sensitivity;})

      return result;
}
function getBGTarget(time): float{
      // look up what the ISF value should be based on the current time the function is being called
      // Lets say its 3am
      // maybe use an array of object, iterate over the array, look at each position , and if the current time
      // falls between the start and end field, return that value.

      const result = bglist.filter(bg => {
            if( bg.start >= time && bg.end < time) return bg.high;
            return 6
      })

      return result;
}
function getCarbRatio(time): float{
      // look up what the ISF value should be based on the current time the function is being called
      // Lets say its 3am
      // maybe use an array of object, iterate over the array, look at each position , and if the current time
      // falls between the start and end field, return that value.

      const result = carbRatioList.filter(ratio => {
            if( ratio.start >= time && ratio.end < time) return ratio.ratio;
            return 6
      })

      return result;
}

function getBgCorrection(bg, bgtarget, ISF){
      // This means the user's blood sugar is low
      if (bg < bgTarget) return 0;

      // This means the user's blood sugar is high
      return ISF * (bg - bgTarget);
}

function getActiveInsulin(time): float{
      // in order for this function to work the user account must have previous saved history.
      // because the bolusWizard function will adjust how much insulin must be given for a high blood sugar correction
      // based on remaining effects of boluses given in the past. If there has been no bolus in the last 2hours (if the users
      // insulin last that long based on what the user has saved), then any bolus that has been mae with in the those two hours
      // must be considered in the correction.

      let inThePast = Date.now() - ACTIVE_INSULIN_TIME;
      let remaingInsulinInEffect = 0;

      bolusHistory.filter(history =>{
            if(history.bolusTime > inThePast) {
                  // this is the remaing effects of the most recent bolus
                  const perecentageOfRemainingInsulinEffects = ( ACTIVE_INSULIN_TIME - (history.bolusTime - inThePast)) / ACTIVE_INSULIN_TIME;
                  
                  remaingInsulinInEffect = remaingInsulinInEffect + (perecentageOfRemainingInsulinEffects * hitsory.bolus);
            }
      })

      return remaingInsulinInEffect;
}
```
