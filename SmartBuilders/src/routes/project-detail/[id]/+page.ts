import type { PageLoad } from './$types';

export const load: PageLoad = async ({ params, fetch }) => {
  const projectId = params.id;

  try {
    // Fetch the project data from your Go backend API
    const response = await fetch(`http://13.49.46.226:8080/api/v1/projects/getByID/${projectId}`);

    if (!response.ok) {
      throw new Error(`Could not fetch project: ${response.statusText}`);
    }

    const project = await response.json();

    // The returned object will be available in +page.svelte as the `data` prop.
    return {
      project: project,
      status: response.status
    };
  } catch (error) {
    console.error('Failed to load project:', error);
    // Return the error to be displayed on the page
    return {
      project: null,
      status: 500,
      error: 'Project not found or an error occurred while fetching.'
    };
  }
};
