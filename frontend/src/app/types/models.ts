export interface DefaultServerResponse {
  status: 'ok' | 'error';
}

export interface ErrorServerResponse extends DefaultServerResponse {
  error_detail: string;
}
