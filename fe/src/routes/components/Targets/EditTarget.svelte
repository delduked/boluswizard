<script lang="ts">
	import Alert from './../Alert.svelte';
	import TargetRow from './TargetRow.svelte';
	import { onMount } from 'svelte';
	import { getTargets, saveTargets } from '../../utils/client';
	import type { Target } from '../../utils/types';

	let rows: Target[];
	let promise = null;

	function addTarget() {
		let row: Target = {
			Start: '00h00m',
			End: '00h00m',
			High: 6.6,
			Low: 6.6
		};
		rows = [...rows, row]; // instead of rows.push(row)
	}

	async function save() {
		try {
			if (rows) {
				const res = await saveTargets(rows)
					.then(data => data)
					.catch(err => {
						throw err;
					});
				if (res.Status == 200) {
					return true;
				} else {
					throw Error('Error saving Target rows');
				}
			} else {
				throw Error('No rows found');
			}
		} catch (error) {
			return error;
		}
	}

	function handleClick() {
		promise = save();

		setTimeout(() => {
			promise = null;
		}, 2000);
	}

	onMount(() => {
		getTargets().then((data) => (rows = data.Data));
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
					<TargetRow bind:rows />
				</tbody>
			</table>
		</div>
		<div class="flex justify-between items-baseline mt-4">
			<div class="join m-3">
				<button
					on:click|preventDefault={addTarget}
					class="btn btn-active btn-secondary btn-sm md:btn-md"
					style="border-top-right-radius: 0;border-bottom-right-radius: 0;">Add Target</button
				>
				<!-- svelte-ignore missing-declaration -->
				<button
					on:click|preventDefault={handleClick}
					class="btn btn-active btn-primary btn-sm md:btn-md"
					style="border-top-left-radius: 0;border-bottom-left-radius: 0;">Save Targets</button
				>
			</div>
		</div>
		<Alert message={'Saving Targets'} {promise} />
		<p class="ml-3 mt-1">Press ESC key or click on ✕ button to close</p>
	</form>
</dialog>
