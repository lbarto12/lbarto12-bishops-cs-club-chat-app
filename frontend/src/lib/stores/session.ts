import type { Session } from "$lib/models/session/Session";
import { writable, type Writable } from "svelte/store";

export const session: Writable<Session> = writable<Session>({
    name: ""
});