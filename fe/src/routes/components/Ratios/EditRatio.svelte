<script lang="ts">
	import RatioRow from './RatioRow.svelte';
	import { onMount } from 'svelte';
	import { get } from "../../utils/client";;
	import type { Ratio } from "../../utils/types";
	let rows: Ratio[];

	onMount(() => {
		try {
			get<Ratio[]>('Ratios').then((data) => (rows = data.Data));
		} catch (error) {
			console.log(error);
		}
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
					<RatioRow {rows}/>
				</tbody>
			</table>
		</div>
		<div class="flex justify-between items-baseline mt-4">
			<p class="py-4 ml-3">Press ESC key or click on ✕ button to close</p>
			<button class="btn btn-active btn-primary">Add Target</button>
		</div>
	</form>
</dialog>
