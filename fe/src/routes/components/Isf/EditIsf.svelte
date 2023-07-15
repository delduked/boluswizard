<script lang="ts">
	import { onMount } from 'svelte';
	import { get, getISFs, post, saveISFs } from '../../utils/client';
	import IsfRow from './IsfRow.svelte';
	import type { Isf, iResponse } from '../../utils/types';

	let rows: Isf[];

	function addIsf() {
		let row: Isf = {
			Start: '00h00m',
			End: '00h00m',
			Sensitivity: 1.1
		};
		rows = [...rows, row]; // instead of rows.push(row)
	}

	const handleClick = async () => {
		saveISFs(rows)
	};

	onMount(() => {
		getISFs().then(data => rows = data.Data)
	});
</script>

<button class="btn bg-neutral border-0" onclick="my_modal_isf.showModal()">ISF</button>
<dialog id="my_modal_isf" class="modal">
	<form method="dialog" class="modal-box w-11/12 max-w-5xl">
		<button class="btn btn-sm btn-circle btn-ghost absolute right-2 top-2">✕</button>
		<h3 class="font-bold text-lg ml-3">Edit Insulin Sensitivity Factor!</h3>
		<div class=" overflow-x-auto">
			<table class="table table-zebra table-sm mt-4">
				<!-- head -->
				<thead>
					<tr>
						<th>Start</th>
						<th>End</th>
						<th>Sensitivity</th>
					</tr>
				</thead>
				<tbody>
					<IsfRow bind:rows />
				</tbody>
			</table>
		</div>
		<div class="flex justify-between items-baseline mt-4">
			<div class="join m-3">
				<button on:click|preventDefault={addIsf} class="btn btn-active btn-secondary btn-sm md:btn-md" style="border-top-right-radius: 0;border-bottom-right-radius: 0;">Add ISF</button>
				<button on:click|preventDefault={handleClick} class="btn btn-active btn-primary btn-sm md:btn-md" style="border-top-left-radius: 0;border-bottom-left-radius: 0;">Save ISFs</button>
			</div>
		</div>
		<p class="ml-3 mt-1">Press ESC key or click on ✕ button to close</p>
		<div/>
	</form>
</dialog>
