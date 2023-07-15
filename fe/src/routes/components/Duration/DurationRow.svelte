<script lang="ts">
	export let row: duration;
	import type { duration } from '../../utils/types';
</script>

{#if row}
	<tr>
		<input
			type="text"
			bind:value={row.Duration}
			pattern="^([01]?[0-9]|2[0-3])h[0-5][0-9]m$"
			placeholder="00h00m"
			class="input input-bordered input-sm w-24 max-w-xs"
		/>
	</tr>
{:else}
	<tr>
		<td>00h00m</td>
	</tr>
{/if}
