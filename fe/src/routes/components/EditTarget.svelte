<script lang="ts">
  import IMask from 'imask';
  import { onMount } from 'svelte';

</script>

<button class="btn bg-neutral border-0" onclick="my_modal.showModal()">Target</button>
<dialog id="my_modal" class="modal">
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
					<!-- row 1 -->
					<tr>
						<td>00h00m</td>
						<td>05h59m</td>
						<td>6.6</td>
						<td>5.4</td>
					</tr>
					<!-- row 2 -->
					<tr>
						<td>06h00m</td>
						<td>11h59m</td>
						<td>6.4</td>
						<td>6.1</td>
					</tr>
					<tr>
						<td>12h00m</td>
						<td>15h59m</td>
						<td>6.5</td>
						<td>5.8</td>
					</tr>
					<tr>
						<td>16h00m</td>
						<td>23h59m</td>
						<td>7.0</td>
						<td>6.0</td>
					</tr>
					<tr>
						<td
							><input
								type="text"
                                pattern="^([01]?[0-2]|2[0-3])h[0-5][0-9]m$" 
								placeholder="00h00m"
								class="input input-bordered input-xs w-24 max-w-xs"
							/>
						</td>
						<td
							><input
								type="text"
                                pattern="^([01]?[0-2]|2[0-3])h[0-5][0-9]m$" 
								placeholder="00h00m"
								class="input input-bordered input-xs w-24 max-w-xs"
							/>
						</td>
						<td
							><input
								type="text"
                                pattern="^4\.[0-9]+|^[5-7](\.[0-9]*)?|^8\.0*$" 
								placeholder="0.0"
								class="input input-bordered input-xs w-24 max-w-xs"
							/>
						</td>
						<td
							><input
								type="text"
                                pattern="^4\.[0-9]+|^[5-7](\.[0-9]*)?|^8\.0*$" 
								placeholder="0.0"
								class="input input-bordered input-xs w-24 max-w-xs"
							/>
						</td>
					</tr>
					<tr>
						<td
							><input
								type="text"
                                pattern="^([01]?[0-2]|2[0-3])h[0-5][0-9]m$" 
								placeholder="00h00m"
								class="input input-bordered input-xs w-24 max-w-xs"
							/>
						</td>
						<td
							><input
								type="text"
                                pattern="^([01]?[0-2]|2[0-3])h[0-5][0-9]m$" 
								placeholder="00h00m"
								class="input input-bordered input-xs w-24 max-w-xs"
							/>
						</td>
						<td
							><input
								type="text"
                                pattern="^4\.[0-9]+|^[5-7](\.[0-9]*)?|^8\.0*$" 
								placeholder="0.0"
								class="input input-bordered input-xs w-24 max-w-xs"
							/>
						</td>
						<td
							><input
								type="text"
                                pattern="^4\.[0-9]+|^[5-7](\.[0-9]*)?|^8\.0*$" 
								placeholder="0.0"
								class="input input-bordered input-xs w-24 max-w-xs"
							/>
						</td>
					</tr>
				</tbody>
			</table>
		</div>
		<div class="flex justify-between items-baseline mt-4">
			<p class="py-4 ml-3">Press ESC key or click on ✕ button to close</p>
			<button class="btn btn-active btn-primary">Add Target</button>
		</div>
	</form>
</dialog>
