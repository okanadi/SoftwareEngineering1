
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
		RouteId(): "/" | "/company" | "/company/projects" | "/create-project" | "/customer" | "/customer/projects" | "/employee" | "/employee/projects" | "/project-detail" | "/project-detail/[id]" | "/projects" | "/signin" | "/signup";
		RouteParams(): {
			"/project-detail/[id]": { id: string }
		};
		LayoutParams(): {
			"/": { id?: string };
			"/company": Record<string, never>;
			"/company/projects": Record<string, never>;
			"/create-project": Record<string, never>;
			"/customer": Record<string, never>;
			"/customer/projects": Record<string, never>;
			"/employee": Record<string, never>;
			"/employee/projects": Record<string, never>;
			"/project-detail": { id?: string };
			"/project-detail/[id]": { id: string };
			"/projects": Record<string, never>;
			"/signin": Record<string, never>;
			"/signup": Record<string, never>
		};
		Pathname(): "/" | "/company" | "/company/" | "/company/projects" | "/company/projects/" | "/create-project" | "/create-project/" | "/customer" | "/customer/" | "/customer/projects" | "/customer/projects/" | "/employee" | "/employee/" | "/employee/projects" | "/employee/projects/" | "/project-detail" | "/project-detail/" | `/project-detail/${string}` & {} | `/project-detail/${string}/` & {} | "/projects" | "/projects/" | "/signin" | "/signin/" | "/signup" | "/signup/";
		ResolvedPathname(): `${"" | `/${string}`}${ReturnType<AppTypes['Pathname']>}`;
		Asset(): "/assets/SmartBuildersBackground.webp" | "/robots.txt" | string & {};
	}
}