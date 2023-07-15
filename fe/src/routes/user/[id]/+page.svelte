<script lang="ts">
	import Nav from '../../components/nav.svelte';
	import Stats from '../../components/Stats.svelte';
	import UserInput from '../../components/UserInput.svelte';

	export let data;
	console.log(data)
</script>

{#if data.auth}
	<Nav />
	<div class="w-full flex justify-center items-center flex-wrap lg:m-4 sm:m-2">
		<UserInput />
		<Stats />
	</div>
	<div class="artboard phone-2">
	</div>
{:else}
	<h1>{data.auth}</h1>
{/if}
