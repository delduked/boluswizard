<script>
	import { browser } from '$app/environment'; // (was '$app/env' in a pre 1.0 SvelteKit version)

	import Nav from '../../components/nav.svelte';
	import Stats from '../../components/Stats.svelte';
	import UserInput from '../../components/UserInput.svelte';
	import { onMount } from 'svelte';

	let chartData = [20, 100, 50, 12, 20, 130, 45];
	let Labels = ['Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday', 'Saturday', 'Sunday'];
	let ctx;
	let chartCanvas;

	
	onMount(async (promise) => {
		const { Chart } = await import('chart.js');
		ctx = await chartCanvas.getContext('2d');
		var chart = new Chart(ctx, {
			type: 'bar',
			data: {
				labels: Labels,
				datasets: [
					{
						label: 'Unit Sales',
						data: chartData
					}
				]
			}
		});
	});
</script>

<Nav />
<div class="w-full flex justify-center items-center flex-wrap lg:m-4 sm:m-2">
	<UserInput />
	
	<Stats />
</div>
<div class="artboard phone-2">
	<canvas bind:this={chartCanvas} />
</div>
