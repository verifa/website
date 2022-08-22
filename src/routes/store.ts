
import { writable } from "svelte/store";
import type { Writable } from "svelte/store";

export const headerVisible: Writable<boolean> = writable(false);