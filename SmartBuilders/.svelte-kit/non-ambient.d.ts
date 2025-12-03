
// this file is generated — do not edit it


declare module "svelte/elements" {
	export interface HTMLAttributes<T> {
		'data-sveltekit-keepfocus'?: true | '' | 'off' | undefined | null;
		'data-sveltekit-noscroll'?: true | '' | 'off' | undefined | null;
		'data-sveltekit-preload-code'?:
			| true
			| ''
			| 'eager'
			| 'viewport'
			| 'hover'
			| 'tap'
			| 'off'
			| undefined
			| null;
		'data-sveltekit-preload-data'?: true | '' | 'hover' | 'tap' | 'off' | undefined | null;
		'data-sveltekit-reload'?: true | '' | 'off' | undefined | null;
		'data-sveltekit-replacestate'?: true | '' | 'off' | undefined | null;
	}
}

export {};


declare module "$app/types" {
	export interface AppTypes {
		RouteId(): "/" | "/company" | "/company/projects" | "/company/projects/project-detail" | "/company/projects/project-detail/edit" | "/create-project" | "/customer" | "/customer/projects" | "/customer/projects/project-detail" | "/customer/projects/project-detail/edit" | "/employee" | "/employee/projects" | "/employee/projects/project-detail" | "/employee/projects/project-detail/edit" | "/signin" | "/signup";
		RouteParams(): {
			
		};
		LayoutParams(): {
			"/": Record<string, never>;
			"/company": Record<string, never>;
			"/company/projects": Record<string, never>;
			"/company/projects/project-detail": Record<string, never>;
			"/company/projects/project-detail/edit": Record<string, never>;
			"/create-project": Record<string, never>;
			"/customer": Record<string, never>;
			"/customer/projects": Record<string, never>;
			"/customer/projects/project-detail": Record<string, never>;
			"/customer/projects/project-detail/edit": Record<string, never>;
			"/employee": Record<string, never>;
			"/employee/projects": Record<string, never>;
			"/employee/projects/project-detail": Record<string, never>;
			"/employee/projects/project-detail/edit": Record<string, never>;
			"/signin": Record<string, never>;
			"/signup": Record<string, never>
		};
		Pathname(): "/" | "/company" | "/company/" | "/company/projects" | "/company/projects/" | "/company/projects/project-detail" | "/company/projects/project-detail/" | "/company/projects/project-detail/edit" | "/company/projects/project-detail/edit/" | "/create-project" | "/create-project/" | "/customer" | "/customer/" | "/customer/projects" | "/customer/projects/" | "/customer/projects/project-detail" | "/customer/projects/project-detail/" | "/customer/projects/project-detail/edit" | "/customer/projects/project-detail/edit/" | "/employee" | "/employee/" | "/employee/projects" | "/employee/projects/" | "/employee/projects/project-detail" | "/employee/projects/project-detail/" | "/employee/projects/project-detail/edit" | "/employee/projects/project-detail/edit/" | "/signin" | "/signin/" | "/signup" | "/signup/";
		ResolvedPathname(): `${"" | `/${string}`}${ReturnType<AppTypes['Pathname']>}`;
		Asset(): "/robots.txt" | string & {};
	}
}