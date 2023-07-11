<script lang="ts">
    export let row: string;
</script>

{#if row}
    <tr>
        <td>
            <input
                type="text"
                value={row}
                pattern="^([01]?[0-9]|2[0-3])h[0-5][0-9]m$" 
                placeholder="00h00m"
                class="input input-bordered input-xs w-24 max-w-xs"
            />
        </td>
    </tr>
{:else}
    <tr>
        <td>00h00m</td>
    </tr>
{/if}

