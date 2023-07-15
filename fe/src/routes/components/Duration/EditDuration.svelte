<script lang="ts">
	import DurationRow from './DurationRow.svelte';
	import { onMount } from 'svelte';
	import { getDuration, saveDuration } from '../../utils/client';
	import type { duration } from '../../utils/types';
	let row: duration;

	const handleClick = async () => {
		saveDuration(row)
	};

	onMount(() => {
		getDuration().then(data => row = data.Data)
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
			<button on:click={handleClick} class="btn btn-active btn-primary m-2 btn-sm md:btn-md"
				>Save Duration</button
			>
		</div>
		<p class="ml-3 mt-1">Press ESC key or click on ✕ button to close</p>

	</form>
</dialog>
