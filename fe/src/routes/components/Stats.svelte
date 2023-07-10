<script lang="ts">
	import { onMount } from 'svelte';
	import { post, get } from '../utils/client';
	import type { duration, iTarget, iRatio, iIsf } from '../utils/types';

	let Isf;
	let Duration;
	let Target;
	let Ratio;

	onMount(() => {
		try {
			get<duration>('Duration').then((data) => (Duration = data.Duration));
			get<iIsf>('CurrentISF').then((data) => (Isf = data));
			get<iTarget>('CurrentTarget').then((data) => (Target = data.Low + '-' + data.High));
		} catch (error) {
			console.log(error);
		}
	});
</script>

<div class="stats shadow flex-col justify-center">
	<div class="stat place-items-center">
		<div class="stat-title">Bolus</div>
		<div class="stat-value text-secondary">N/A</div>
		<div class="stat-desc text-secondary">units</div>
	</div>
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
			<div class="stat-value text-info">{Isf}</div>
		{:else}
			<div class="stat-value text-info">n/a</div>
		{/if}

		<div class="stat-desc text-info">unit/mmol</div>
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
</div>
