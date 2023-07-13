<script lang="ts">
    import type { Target } from "../../utils/types";
    export let rows: Target[]
</script>

{#if rows}
    {#each rows as row}
    <tr>
        <td>
            <input
                bind:value={row.Start}
                type="text"
                pattern="^([01]?[0-9]|2[0-3])h[0-5][0-9]m$" 
                placeholder="00h00m"
                class="input input-bordered input-xs w-24 max-w-xs"
            />
        </td>
        <td>
            <input
                type="text"
                bind:value={row.End}
                pattern="^([01]?[0-9]|2[0-3])h[0-5][0-9]m$" 
                placeholder="00h00m"
                class="input input-bordered input-xs w-24 max-w-xs"
            />
        </td>
        <td>
            <input
                type="number"
                bind:value={row.High}
                pattern="^4\.[0-9]+|^[5-7](\.[0-9]*)?|^8\.0*$" 
                placeholder="0.0"
                class="input input-bordered input-xs w-24 max-w-xs"
            />
        </td>
        <td>
            <input
                type="number"
                bind:value={row.Low}
                pattern="^4\.[0-9]+|^[5-7](\.[0-9]*)?|^8\.0*$" 
                placeholder="0.0"
                class="input input-bordered input-xs w-24 max-w-xs"
            />
        </td>
    </tr>
    {/each}
{:else}
    <tr>
        <td>00h00m</td>
        <td>00h00m</td>
        <td>0.0</td>
        <td>0.0</td>
    </tr>
{/if}

