<script lang="ts">
	import DurationRow from './DurationRow.svelte';
	import { onMount } from 'svelte';
	import { get } from "../../utils/client";;
	import type { duration } from "../../utils/types";
	let row: string;

	onMount(() => {
		try {
			get<duration>('Duration').then((data) => (row = data.Data.Duration));
		} catch (error) {
			console.log(error);
		}
	});
</script>

<button class="btn bg-neutral border-0" onclick="my_modal_duration.showModal()">Duration</button>
<dialog id="my_modal_duration" class="modal">
	<form method="dialog" class="modal-box">
		<button class="btn btn-sm btn-circle btn-ghost absolute right-2 top-2">✕</button>
		<h3 class="font-bold text-lg ml-3">Edit Targets!</h3>
		<div class=" overflow-x-auto">
			<table class="table table-zebra table-sm mt-4">
				<!-- head -->
				<thead>
					<tr>
						<th>Duration</th>
					</tr>
				</thead>
				<tbody>
					<DurationRow {row}/>
				</tbody>
			</table>
		</div>
		<div class="flex justify-between items-baseline mt-4">
			<p class="py-4 ml-3">Press ESC key or click on ✕ button to close</p>
		</div>
	</form>
</dialog>
