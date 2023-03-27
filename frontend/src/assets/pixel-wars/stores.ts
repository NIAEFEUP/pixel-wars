import { writable } from "svelte/store";


export const ColorPickerStore = writable(0);

export const TimeoutStore = writable({timeout: new Date(), remainingPixels: 0})