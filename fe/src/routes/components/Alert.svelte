<script>
	import { fly } from 'svelte/transition';
	export let promise = null;
    export let message;
</script>

{#if promise != null}
	{#await promise}
		<p />
	{:then data}
		{#if data == true}
			<div
				in:fly={{ y: -10, duration: 500 }}
				out:fly={{ y: -10, duration: 500 }}
				class="alert alert-success"
			>
				<svg
					xmlns="http://www.w3.org/2000/svg"
					class="stroke-current shrink-0 h-6 w-6"
					fill="none"
					viewBox="0 0 24 24"
					><path
						stroke-linecap="round"
						stroke-linejoin="round"
						stroke-width="2"
						d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"
					/></svg
				>
				<span>{message} Successfull!</span>
			</div>
		{:else}
			<div
				in:fly={{ y: -10, duration: 500 }}
				out:fly={{ y: -10, duration: 500 }}
				class="alert alert-error"
			>
				<svg
					xmlns="http://www.w3.org/2000/svg"
					class="stroke-current shrink-0 h-6 w-6"
					fill="none"
					viewBox="0 0 24 24"
					><path
						stroke-linecap="round"
						stroke-linejoin="round"
						stroke-width="2"
						d="M10 14l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2m7-2a9 9 0 11-18 0 9 9 0 0118 0z"
					/></svg
				>
				<span>{message} failed!</span>
			</div>
		{/if}
	{/await}
{/if}
