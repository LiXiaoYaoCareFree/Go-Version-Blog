/// <reference types="vite/client" />
declare module '*.vue' {
    import type {DefineComponent} from 'vue';
    const component:DefineComponent<{}, {}, any>;

    export default component;
}


import "vue-router"
declare module 'vue-router' {
    interface RouteMeta {
        title: string;
    }
}

export interface EnvMeta extends Record<string, string>{
    ViTE_SERVER_URL: string;
}