<script lang="ts">
	import { calculateCorrection } from "../utils/client";

	let bg: number;
	let carb: number;

	let BgCorrection: number;
	let CarbCorrection: number;
	let ActiveInsulinReduction: number;
	let Bolus: number;

</script>

<div class="flex flex-wrap justify-center items-center">
	<div style="width: 240px;" class="w-80 flex justify-center items-center flex-wrap m-4">
		<div class="join mb-4">
			<input
				bind:value={bg}
				type="number"
				placeholder="Enter BG"
				class="input input-bordered input-sm input-secondary w-full max-w-xs"
				style="border-top-right-radius: 0;border-bottom-right-radius: 0;"
			/>
			<button
				on:click={() =>{
					calculateCorrection(bg,carb).then(data => {
						BgCorrection = data.BgCorrection
						CarbCorrection = data.CarbCorrection
						ActiveInsulinReduction = data.ActiveInsulinReduction
						Bolus = data.Bolus
					})
				}}
				class="btn btn-sm btn-secondary"
				style="border-top-left-radius: 0; border-bottom-left-radius: 0;">CALC</button
			>
		</div>

		<div class="join">
			<input
				bind:value={carb}
				type="number"
				placeholder="Enter Carbs"
				class="input input-bordered input-sm input-info w-full max-w-xs"
				style="border-top-right-radius: 0;border-bottom-right-radius: 0;"
			/>
			<button
				class="btn btn-sm btn-info"
				style="border-top-left-radius: 0; border-bottom-left-radius: 0;">SAVE</button
			>
		</div>
	</div>
	<div class="stats bg-primary text-primary-content m-3" style="width: 240px;">
		<div class="stat">
			<div class="stat-title">Correction</div>
			{#if BgCorrection }
			<div class="stat-value" style="border-top-right-radius: 0;border-bottom-right-radius: 0;">{Bolus}</div>
			{:else}
			<div class="stat-value" style="border-top-right-radius: 0;border-bottom-right-radius: 0;">n/a</div>
			{/if}
		</div>
	</div>
</div>
<div class="w-full flex justify-center items-center mt-3 mb-3 ml-2 mr-2 flex-wrap">
	<div class="join m-2">
		<div class="badge badge-neutral">Bg Correction</div>
		{#if BgCorrection }
			<div class="badge badge-accent" style="border-top-left-radius: 0; border-bottom-left-radius: 0;">{BgCorrection}</div>
		{:else}
			<div class="badge badge-accent" style="border-top-left-radius: 0; border-bottom-left-radius: 0;">n/a</div>
		{/if}
	</div>
	<div class="join m-2">
		<div class="badge badge-neutral">Carb Correction</div>
		{#if CarbCorrection }
			<div class="badge badge-accent" style="border-top-left-radius: 0; border-bottom-left-radius: 0;">{CarbCorrection}</div>
		{:else}
			<div class="badge badge-accent" style="border-top-left-radius: 0; border-bottom-left-radius: 0;">n/a</div>
		{/if}
	</div>
	<div class="join m-2">
		<div class="badge badge-neutral">Active Insulin Reduction</div>
		{#if ActiveInsulinReduction }
		<div class="badge badge-accent" style="border-top-left-radius: 0; border-bottom-left-radius: 0;">{ActiveInsulinReduction}</div>
		{:else}
		<div class="badge badge-accent" style="border-top-left-radius: 0; border-bottom-left-radius: 0;">n/a</div>
		{/if}
	</div>
</div>
