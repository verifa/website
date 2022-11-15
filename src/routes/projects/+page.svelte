<script lang="ts">
	import { seo } from '$lib/seo/store';
	import Project, { type ProjectConfig } from './project.svelte';

	const projects: ProjectConfig[] = [
		{
			name: 'Terraplate',
			githubRepo: 'verifa/terraplate'
		},
		{
			name: 'Artifactory Retention',
			githubRepo: 'verifa/rt-retention'
		},
		{
			name: 'Learning Centre',
			githubRepo: 'verifa/learn'
		},
		{
			name: 'Website',
			githubRepo: 'verifa/website'
		}
	];

	seo.reset();
	$seo.title = 'Projects';
	$seo.description = 'Collection of open source projects by Verifa.';
</script>

<section>
	<h1>Projects</h1>
	<h3>Here's a collection of open source projects created by Verifa.</h3>
</section>
<section>
	<div class="grid grid-cols-1 md:grid-cols-2 gap-12">
		{#each projects as project}
			<Project {project} />
		{/each}
	</div>
</section>
