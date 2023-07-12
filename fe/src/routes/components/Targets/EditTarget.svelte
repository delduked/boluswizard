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
			await post<iResponse<Target[]>, Target[]>('Targets', rows).catch((err) => {
				throw err;
			});
		} catch (error) {
			console.log(error);
		}
	};
	let promise
	function handleClick() {
		promise = saveTargets();
	}

	onMount(() => {
		try {
			get<Target[]>('Targets').then((data) => (rows = data.Data));
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
			<button on:click|preventDefault={addTarget} class="btn btn-active btn-primary"
				>Add Target</button
			>
			<button on:click|preventDefault={handleClick} class="btn btn-active btn-primary"
				>Save Target</button
			>
			{#await promise}
				<span class="loading loading-spinner loading-sm"></span>
			{:then data}
				<div class="alert alert-success">
					<svg xmlns="http://www.w3.org/2000/svg" class="stroke-current shrink-0 h-6 w-6" fill="none" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" /></svg>
					<span>Targets Saved!</span>
				</div>
			{:catch error}
				<div class="alert alert-error">
					<svg xmlns="http://www.w3.org/2000/svg" class="stroke-current shrink-0 h-6 w-6" fill="none" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 14l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2m7-2a9 9 0 11-18 0 9 9 0 0118 0z" /></svg>
					<span>{error}</span>
				</div>
			{/await}

		</div>
	</form>
</dialog>
