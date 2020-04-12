export type CSRFHeaderName = 'X-CSRF-Token';
export type EncodedHeader = string;
export interface CSRFHeader { name: CSRFHeaderName; value: string; }

export class InvalidCSRFHeader extends Error {}
export class NoPublicKey extends Error {}
