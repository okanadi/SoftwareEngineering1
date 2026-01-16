import type { PageLoad } from './$types';

export const load: PageLoad = async ({ params, fetch }) => {
  const projectId = params.id;

  try {
    // Fetch project steps from the backend
    const response = await fetch(`http://13.49.46.226:8080/api/v1/project-steps/getAllByProjectID/${projectId}`);

    if (!response.ok) {
      throw new Error(`Could not fetch project steps: ${response.statusText}`);
    }

    const steps = await response.json();

    return {
      projectId,
      steps: Array.isArray(steps) ? steps : [],
      status: response.status
    };
  } catch (error) {
    console.error('Failed to load project steps:', error);
    return {
      projectId,
      steps: [],
      status: 500,
      error: 'Could not fetch project steps or no steps found.'
    };
  }
};
