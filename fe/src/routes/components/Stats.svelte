<script lang="ts">
	import { onMount } from 'svelte';
	import { get } from '../utils/client';
	import type { duration, iTarget } from '../utils/types';

	let Isf;
	let Duration;
	let Target;
	let Ratio;
	onMount(() => {
		try {
			get<duration>('Duration').then((data) => (Duration = data.Data.Duration));
			get<number>('CurrentISF').then((data) => (Isf = data.Data));
			get<iTarget>('CurrentTarget').then((data) => (Target = data.Data.Low + '-' + data.Data.High));
			get<number>('CurrentRatio').then((data) => Ratio = data.Data);
		} catch (error) {
			console.log(error);
		}
	});
</script>

<div class="stats shadow flex-col justify-center">
	<div class="stat place-items-center">
		<div class="stat-title">Target</div>
		{#if Target}
			<div class="stat-value">{Target}</div>
		{:else}
			<div class="stat-value">n/a</div>
		{/if}
		<div class="stat-desc">mmols</div>
	</div>

	<div class="stat place-items-center">
		<div class="stat-title">ISF</div>
		{#if Isf}
			<div class="stat-value text-secondary">{Isf}</div>
		{:else}
			<div class="stat-value text-secondary">n/a</div>
		{/if}

		<div class="stat-desc text-secondary">unit/mmol</div>
	</div>

	<div class="stat place-items-center">
		<div class="stat-title">Duration</div>
		{#if Duration}
			<div class="stat-value">{Duration}</div>
		{:else}
			<div class="stat-value">n/a</div>
		{/if}
		<div class="stat-desc">hours</div>
	</div>
	<div class="stat place-items-center">
		<div class="stat-title">Ratio</div>
		{#if Ratio}
			<div class="stat-value text-info">{Ratio}</div>
		{:else}
			<div class="stat-value text-info">n/a</div>
		{/if}
		<div class="stat-desc">units/carbs</div>
	</div>
</div>
