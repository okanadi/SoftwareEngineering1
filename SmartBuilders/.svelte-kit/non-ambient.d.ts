
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
		RouteId(): "/" | "/create-project" | "/project-detail" | "/project-detail/[id]" | "/project-detail/[id]/steps" | "/projects" | "/signin" | "/signup";
		RouteParams(): {
			"/project-detail/[id]": { id: string };
			"/project-detail/[id]/steps": { id: string }
		};
		LayoutParams(): {
			"/": { id?: string };
			"/create-project": Record<string, never>;
			"/project-detail": { id?: string };
			"/project-detail/[id]": { id: string };
			"/project-detail/[id]/steps": { id: string };
			"/projects": Record<string, never>;
			"/signin": Record<string, never>;
			"/signup": Record<string, never>
		};
		Pathname(): "/" | "/create-project" | "/create-project/" | "/project-detail" | "/project-detail/" | `/project-detail/${string}` & {} | `/project-detail/${string}/` & {} | `/project-detail/${string}/steps` & {} | `/project-detail/${string}/steps/` & {} | "/projects" | "/projects/" | "/signin" | "/signin/" | "/signup" | "/signup/";
		ResolvedPathname(): `${"" | `/${string}`}${ReturnType<AppTypes['Pathname']>}`;
		Asset(): "/assets/SmartBuildersBackground.webp" | "/assets/SmartBuildersBackgroundGradient.jpg" | "/robots.txt" | string & {};
	}
}