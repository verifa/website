<script lang="ts" context="module">
	export interface ProjectConfig {
		name: string;
		githubRepo: string;
	}
</script>

<script lang="ts">
	import { onMount } from 'svelte';

	interface RepoData {
		description: string;
		html_url: string;
		homepage: string;
		stargazers_count: number;
	}

	export let project: ProjectConfig;

	let repo: RepoData;
	let err: Error;

	onMount(() => {
		fetch(`https://api.github.com/repos/${project.githubRepo}`, {
			headers: {
				Accept: 'application/vnd.github+json'
			}
		})
			.then((resp) => {
				resp
					.json()
					.then((data) => {
						repo = data;
					})
					.catch((error) => {
						err = error;
					});
			})
			.catch((error) => {
				err = error;
			});
	});
</script>

<div class="border-4 border-v-black border-b-8 border-r-8 p-4 h-full flex flex-col">
	{#if repo}
		<a class="group" target="_blank" href={repo.homepage || repo.html_url}>
			<h3 class="mb-0 group-hover:text-v-gray">{project.name}</h3>
			<p class="group-hover:text-v-gray">{repo.description}</p>
		</a>
		<div class="mt-auto">
			<a target="_blank" href={`https://github.com/${project.githubRepo}`}>
				<div class="group flex items-center space-x-4 hover:cursor-pointer ">
					<svg
						xmlns="http://www.w3.org/2000/svg"
						fill="currentColor"
						class="w-12 h-12 group-hover:text-v-gray"
						viewBox="0 0 16 16"
					>
						<path
							d="M8 0C3.58 0 0 3.58 0 8c0 3.54 2.29 6.53 5.47 7.59.4.07.55-.17.55-.38 0-.19-.01-.82-.01-1.49-2.01.37-2.53-.49-2.69-.94-.09-.23-.48-.94-.82-1.13-.28-.15-.68-.52-.01-.53.63-.01 1.08.58 1.23.82.72 1.21 1.87.87 2.33.66.07-.52.28-.87.51-1.07-1.78-.2-3.64-.89-3.64-3.95 0-.87.31-1.59.82-2.15-.08-.2-.36-1.02.08-2.12 0 0 .67-.21 2.2.82.64-.18 1.32-.27 2-.27.68 0 1.36.09 2 .27 1.53-1.04 2.2-.82 2.2-.82.44 1.1.16 1.92.08 2.12.51.56.82 1.27.82 2.15 0 3.07-1.87 3.75-3.65 3.95.29.25.54.73.54 1.48 0 1.07-.01 1.93-.01 2.2 0 .21.15.46.55.38A8.012 8.012 0 0 0 16 8c0-4.42-3.58-8-8-8z"
						/>
					</svg>
					<div class="flex flex-col">
						<span class="font-bold group-hover:text-v-gray">{project.githubRepo}</span>
						<div class="flex items-center gap-x-2">
							<svg
								xmlns="http://www.w3.org/2000/svg"
								fill="none"
								viewBox="0 0 24 24"
								stroke-width="1.5"
								stroke="currentColor"
								class="w-6 h-6 group-hover:text-v-gray"
							>
								<path
									stroke-linecap="round"
									stroke-linejoin="round"
									d="M9.813 15.904L9 18.75l-.813-2.846a4.5 4.5 0 00-3.09-3.09L2.25 12l2.846-.813a4.5 4.5 0 003.09-3.09L9 5.25l.813 2.846a4.5 4.5 0 003.09 3.09L15.75 12l-2.846.813a4.5 4.5 0 00-3.09 3.09zM18.259 8.715L18 9.75l-.259-1.035a3.375 3.375 0 00-2.455-2.456L14.25 6l1.036-.259a3.375 3.375 0 002.455-2.456L18 2.25l.259 1.035a3.375 3.375 0 002.456 2.456L21.75 6l-1.035.259a3.375 3.375 0 00-2.456 2.456zM16.894 20.567L16.5 21.75l-.394-1.183a2.25 2.25 0 00-1.423-1.423L13.5 18.75l1.183-.394a2.25 2.25 0 001.423-1.423l.394-1.183.394 1.183a2.25 2.25 0 001.423 1.423l1.183.394-1.183.394a2.25 2.25 0 00-1.423 1.423z"
								/>
							</svg>

							<span class="group-hover:text-v-gray">
								{repo.stargazers_count}
							</span>
						</div>
					</div>
				</div>
			</a>
		</div>
	{:else if err}
		{err}
	{/if}
</div>
