<script lang="ts">
	import Nav from '../../components/nav.svelte';
	import Stats from '../../components/Stats.svelte';
	import UserInput from '../../components/UserInput.svelte';
	import { fly, fade } from 'svelte/transition';

	export let data;
	console.log(data);
</script>

{#if data.auth}
	<div>
		<Nav />
		<div class="w-full flex justify-center items-center flex-wrap lg:m-4 sm:m-2">
			<UserInput />
			<Stats />
		</div>
		<div class="artboard phone-2" />
	</div>
{:else}
	<div
		class="hero min-h-screen"
		style="background-image: url(https://boluswizard.io/assets/loginpage/bg.gif);"
	>
		<div class="hero-overlay" />
		<div class="hero-content text-center text-neutral-content">
			<div class="max-w-md text-white">
				<h1 class="mb-5 text-5xl font-bold">Account not accessible</h1>
				<h3 class="m-4">Please either create an account or signin with an existing account.</h3>
				<a class="btn btn-primary" href="/signin">Sign In</a>
				<a class="btn btn-secondary ml-4" href="/signup">Sign Up</a>
			</div>
		</div>
	</div>
{/if}
