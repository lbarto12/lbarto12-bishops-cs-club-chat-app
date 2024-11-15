import { writable, type Writable } from "svelte/store";

export const size: Writable<number> = writable<number>(0);