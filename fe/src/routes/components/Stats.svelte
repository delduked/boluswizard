<script lang="ts">
	import { fade, blur, draw, fly } from 'svelte/transition';
	import { onMount } from 'svelte';
	import {
		get,
		getDuration,
		getCurrentISF,
		getCurrentTarget,
		getCurrentRatio
	} from '../utils/client';
	import type { duration, iTarget } from '../utils/types';

	let Isf;
	let Duration;
	let Target;
	let Ratio;
	onMount(() => {
		try {
			getDuration().then((data) => (Duration = data.Data.Duration));
			getCurrentISF().then((data) => (Isf = data.Data.Sensitivity));
			getCurrentTarget().then((data) => (Target = data.Data.Low + '-' + data.Data.High));
			getCurrentRatio().then((data) => (Ratio = data.Data.Ratio));
		} catch (error) {
			console.log(error);
		}
	});
</script>

<div class="stats shadow flex-col justify-center">
	<div class="stat place-items-center">
		<div class="stat-title">Target</div>
		{#if Target}
			<div transition:fly={{y: -5, duration: 300, delay:300}} class="stat-value">{Target}</div>
		{:else}
			<div class="stat-value"></div>
		{/if}
		<div class="stat-desc">mmols</div>
	</div>

	<div class="stat place-items-center">
		<div class="stat-title">ISF</div>
		{#if Isf}
			<div transition:fly={{y: -5, duration: 300, delay:300}} class="stat-value text-secondary">
				{Isf}
			</div>
		{:else}
			<div class="stat-value text-secondary">n/a</div>
		{/if}

		<div class="stat-desc text-secondary">unit/mmol</div>
	</div>

	<div class="stat place-items-center">
		<div class="stat-title">Duration</div>
		{#if Duration}
			<div transition:fly={{y: -5, duration: 300, delay:300}} class="stat-value">
				{Duration}
			</div>
		{:else}
			<div class="stat-value">n/a</div>
		{/if}
		<div class="stat-desc">hours</div>
	</div>
	<div class="stat place-items-center">
		<div class="stat-title">Ratio</div>
		{#if Ratio}
			<div transition:fly={{y: -5, duration: 300, delay:300}} class="stat-value text-info">
				{Ratio}
			</div>
		{:else}
			<div class="stat-value text-info">n/a</div>
		{/if}
		<div class="stat-desc">units/carbs</div>
	</div>
</div>
