
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
		RouteId(): "/" | "/company" | "/company/projects" | "/company/projects/[id]" | "/company/projects/[id]/edit" | "/create-project" | "/customer" | "/customer/projects" | "/customer/projects/[id]" | "/customer/projects/[id]/edit" | "/employee" | "/employee/projects" | "/employee/projects/[id]" | "/employee/projects/[id]/edit" | "/project-detail" | "/project-detail/[id]" | "/signin" | "/signup";
		RouteParams(): {
			"/company/projects/[id]": { id: string };
			"/company/projects/[id]/edit": { id: string };
			"/customer/projects/[id]": { id: string };
			"/customer/projects/[id]/edit": { id: string };
			"/employee/projects/[id]": { id: string };
			"/employee/projects/[id]/edit": { id: string };
			"/project-detail/[id]": { id: string }
		};
		LayoutParams(): {
			"/": { id?: string };
			"/company": { id?: string };
			"/company/projects": { id?: string };
			"/company/projects/[id]": { id: string };
			"/company/projects/[id]/edit": { id: string };
			"/create-project": Record<string, never>;
			"/customer": { id?: string };
			"/customer/projects": { id?: string };
			"/customer/projects/[id]": { id: string };
			"/customer/projects/[id]/edit": { id: string };
			"/employee": { id?: string };
			"/employee/projects": { id?: string };
			"/employee/projects/[id]": { id: string };
			"/employee/projects/[id]/edit": { id: string };
			"/project-detail": { id?: string };
			"/project-detail/[id]": { id: string };
			"/signin": Record<string, never>;
			"/signup": Record<string, never>
		};
		Pathname(): "/" | "/company" | "/company/" | "/company/projects" | "/company/projects/" | `/company/projects/${string}` & {} | `/company/projects/${string}/` & {} | `/company/projects/${string}/edit` & {} | `/company/projects/${string}/edit/` & {} | "/create-project" | "/create-project/" | "/customer" | "/customer/" | "/customer/projects" | "/customer/projects/" | `/customer/projects/${string}` & {} | `/customer/projects/${string}/` & {} | `/customer/projects/${string}/edit` & {} | `/customer/projects/${string}/edit/` & {} | "/employee" | "/employee/" | "/employee/projects" | "/employee/projects/" | `/employee/projects/${string}` & {} | `/employee/projects/${string}/` & {} | `/employee/projects/${string}/edit` & {} | `/employee/projects/${string}/edit/` & {} | "/project-detail" | "/project-detail/" | `/project-detail/${string}` & {} | `/project-detail/${string}/` & {} | "/signin" | "/signin/" | "/signup" | "/signup/";
		ResolvedPathname(): `${"" | `/${string}`}${ReturnType<AppTypes['Pathname']>}`;
		Asset(): "/assets/SmartBuildersBackground.webp" | "/robots.txt" | string & {};
	}
}