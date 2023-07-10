export interface iRangeData {
	Key: string;
	Bg: number;
	Carbs: number;
	Bolus: number;
	TimeStamp: number;
}

export class RangeData implements iRangeData {
	Key: string;
	Bg: number;
	Carbs: number;
	Bolus: number;
	TimeStamp: number;
}

export interface iCorrectionResponse {
	Status: number;
	Error: any;
	Data: iCorrectionData;
}
export interface iResponse<t> {
	Status: number;
	Error: any;
	Data: t;
}
export interface iCorrectionData {
	BgCorrection: number;
	CarbCorrection: number;
	ActiveInsulinReduction: number;
	Bolus: number;
}
export interface iRangeResponse {
	Status: number;
	Error: any;
	Data: iRangeData[];
}
export interface iRangeData {
	Key: string;
	Bg: number;
	Carbs: number;
	Bolus: number;
	TimeStamp: number;
}

export interface iRatio {
	Key?: string;
	End: string;
	Start: string;
	Ratio: number;
}
export interface iIsf {
	Key?: string;
	End: string;
	Start: string;
	Sensitivity: number;
}
export interface duration {
	Duration: string;
}
export interface iTarget {
	Key?: string;
	End: string;
	Start: string;
	High: number;
	Low: number;
	Sensitivity: number;
}

export interface iCoordinates {
	x: any;
	y: number;
}
export interface iCredentials {
    Username: string;
    Password: string;
}