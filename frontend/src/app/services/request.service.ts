import { Injectable } from '@angular/core';
import {CsrfService} from './classes/csrf/csrf';
import {HttpClient, HttpRequest, HttpResponse} from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class RequestService {
  private readonly csrf: CsrfService = new CsrfService();

  constructor(
    private readonly http: HttpClient
  ) { }

  private static getAPIUrl(endpoint: string): string {
    return `/api/${endpoint}`;
  }

  public async get<T>(endpoint: string): Promise<T> {
    const response = (await this.http.request(new HttpRequest(
      'GET',
      RequestService.getAPIUrl(endpoint)
    )).toPromise()) as HttpResponse<T>;
    const csrfPublicKey = response.headers.get(this.csrf.csrfHeaderName);

    if (!!csrfPublicKey) {
      this.csrf.setPublicKey(csrfPublicKey);
    }

    return response.body;
  }

  public async post<T>(endpoint: string, data: string): Promise<T> {
    return await this.makeRequest(new HttpRequest<any>(
      'POST',
      RequestService.getAPIUrl(endpoint),
      data
    ));
  }

  public async patch<T>(endpoint: string, data: string = ''): Promise<T> {
    return await this.makeRequest(new HttpRequest<any>(
      'PATCH',
      RequestService.getAPIUrl(endpoint),
      data
    ));
  }

  public async delete<T>(endpoint: string, data: string = ''): Promise<T> {
    return await this.makeRequest(new HttpRequest<any>(
      'DELETE',
      RequestService.getAPIUrl(endpoint),
      data
    ));
  }

  private async makeRequest<T>(request: HttpRequest<any>): Promise<T> {
    const csrfHeader = this.csrf.getCSRFHeader();
    request.headers
      .set('Content-Type', 'application/json')
      .set(csrfHeader.name, csrfHeader.value);

    const response = (await this.http.request(request).toPromise()) as HttpResponse<T>;
    this.csrf.setPublicKey(response.headers.get(this.csrf.csrfHeaderName));

    return response.body;
  }
}
