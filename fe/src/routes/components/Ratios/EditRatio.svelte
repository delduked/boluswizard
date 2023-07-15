<script lang="ts">
	import RatioRow from './RatioRow.svelte';
	import { onMount } from 'svelte';
	import { get, getRatios, post, saveRatios } from "../../utils/client";;
	import type { Ratio, iResponse } from "../../utils/types";
	let rows: Ratio[];

	function addRatio() {
		let row: Ratio = {
			Start: "00h00m",
			End: "00h00m",
			Ratio: 0.0,
		}
		rows = [...rows, row]; // instead of rows.push(row)
	}

	const handleClick = () => {
		saveRatios(rows)
	};

	onMount(() => {

		getRatios().then(data => rows = data.Data)

	});
</script>

<button class="btn bg-neutral border-0" onclick="my_modal_ratio.showModal()">Ratio</button>
<dialog id="my_modal_ratio" class="modal">
	<form method="dialog" class="modal-box w-11/12 max-w-5xl">
		<button class="btn btn-sm btn-circle btn-ghost absolute right-2 top-2">✕</button>
		<h3 class="font-bold text-lg ml-3">Edit Ratio!</h3>
		<div class=" overflow-x-auto">
			<table class="table table-zebra table-sm mt-4">
				<!-- head -->
				<thead>
					<tr>
						<th>Start</th>
						<th>End</th>
						<th>Ratio</th>
					</tr>
				</thead>
				<tbody>
					<RatioRow bind:rows />
				</tbody>
			</table>
		</div>
		<div class="flex justify-between items-baseline mt-4">
			<div class="join m-3">
				<button  on:click|preventDefault={addRatio} class="btn btn-active btn-secondary btn-sm md:btn-md" style="border-top-right-radius: 0;border-bottom-right-radius: 0;">Add Ratio</button>
				<button  on:click|preventDefault={handleClick} class="btn btn-active btn-primary btn-sm md:btn-md" style="border-top-left-radius: 0;border-bottom-left-radius: 0;">Save Ratios</button>
			</div>
			<p class="ml-3 mt-1">Press ESC key or click on ✕ button to close</p>
		</div>
	</form>
</dialog>
