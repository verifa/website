<script lang="ts">
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
</script>

<section>
	<h1>Projects</h1>
	<h3>Here's a collection of open source projects Verifa has created or contributed to.</h3>
</section>
<section>
	<div class="grid grid-cols-1 md:grid-cols-2 gap-12">
		{#each projects as project}
			<Project {project} />
		{/each}
	</div>
</section>
