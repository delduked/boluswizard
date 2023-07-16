<script lang="ts">
	import Alert from './../Alert.svelte';

	import DurationRow from './DurationRow.svelte';
	import { onMount } from 'svelte';
	import { getDuration, saveDuration } from '../../utils/client';
	import type { duration } from '../../utils/types';
	let row: duration;
	let promise = null;

	async function save() {
		try {
			if (row) {
				const res = await saveDuration(row)
					.then((data) => data)
					.catch((err) => {
						throw err;
					});
				if (res.Status == 200) {
					return true;
				} else {
					throw Error('Error saving Durection rows');
				}
			} else {
				throw Error('No rows found');
			}
		} catch (error) {
			return error;
		}
	}

	const handleClick = async () => {
		promise = save();

		setTimeout(() => {
			promise = null;
		}, 2000);
	};

	onMount(() => {
		getDuration().then((data) => (row = data.Data));
	});
</script>

<button class="btn bg-neutral border-0" onclick="my_modal_duration.showModal()">Duration</button>
<dialog id="my_modal_duration" class="modal">
	<form method="dialog" class="modal-box">
		<button class="btn btn-sm btn-circle btn-ghost absolute right-2 top-2">✕</button>
		<h3 class="font-bold text-lg ml-3">Edit Duration!</h3>
		<div class=" overflow-x-auto">
			<table class="table table-zebra table-sm mt-4">
				<!-- head -->
				<thead>
					<tr>
						<th>Duration</th>
					</tr>
				</thead>
				<tbody>
					<DurationRow bind:row />
				</tbody>
			</table>
		</div>
		<div class="flex justify-between items-baseline mt-2">
			<button on:click|preventDefault={handleClick} class="btn btn-active btn-primary m-2 btn-sm md:btn-md"
				>Save Duration</button
			>
		</div>
		<Alert message={'Saving Duration'} {promise} />
		<p class="ml-3 mt-1">Press ESC key or click on ✕ button to close</p>
	</form>
</dialog>
