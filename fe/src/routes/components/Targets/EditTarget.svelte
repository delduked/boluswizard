<script lang="ts">
	import TargetRow from './TargetRow.svelte';
	import { onMount } from 'svelte';
	import { get, post } from '../../utils/client';
	import type { Target, iResponse } from '../../utils/types';
	let rows: Target[];

	function addTarget() {
		let row: Target = {
			Start: '00h00m',
			End: '00h00m',
			High: 6.6,
			Low: 6.6
		};
		rows = [...rows, row]; // instead of rows.push(row)
	}

	const saveTargets = async () => {
		try {
			await post<iResponse<Target[]>, Target[]>('Targets', rows)
				.catch((err) => {
					throw err;
				});
		} catch (error) {
			console.log(error);
		}
	};

	onMount(() => {
		try {
			get<Target[]>('Targets')
				.then((data) => (rows = data.Data));
		} catch (error) {
			console.log(error);
		}
	});
</script>

<button class="btn bg-neutral border-0" onclick="my_modal_target.showModal()">Target</button>
<dialog id="my_modal_target" class="modal">
	<form method="dialog" class="modal-box w-11/12 max-w-5xl">
		<button class="btn btn-sm btn-circle btn-ghost absolute right-2 top-2">✕</button>
		<h3 class="font-bold text-lg ml-3">Edit Targets!</h3>
		<div class=" overflow-x-auto">
			<table class="table table-zebra table-sm mt-4">
				<!-- head -->
				<thead>
					<tr>
						<th>Start</th>
						<th>End</th>
						<th>High</th>
						<th>Low</th>
					</tr>
				</thead>
				<tbody>
					<TargetRow {rows} />
				</tbody>
			</table>
		</div>
		<div class="flex justify-between items-baseline mt-4">
			<p class="py-4 ml-3">Press ESC key or click on ✕ button to close</p>
			<div>
				<button on:click|preventDefault={addTarget} class="btn btn-active btn-secondary mr-3">Add Target</button>
				<button on:click|preventDefault={saveTargets} class="btn btn-active btn-primary">Save Targets</button>
			</div>
		</div>
	</form>
</dialog>
