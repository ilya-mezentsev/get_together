import { Injectable } from '@angular/core';
import {CSRFHeader, CSRFHeaderName, EncodedHeader, InvalidCSRFHeader, NoPublicKey} from './types';
import {environment} from '../../../environments/environment';
import {Md5} from 'ts-md5';

@Injectable({
  providedIn: 'root'
})
export class CsrfService {
  private publicKey: string;
  public readonly csrfHeaderName: CSRFHeaderName = 'X-CSRF-Token';

  public setPublicKey(value: EncodedHeader) {
    try {
      this.publicKey = atob(value);
    } catch {
      throw new InvalidCSRFHeader();
    }
  }

  public getCSRFHeader(): CSRFHeader {
    if (!this.publicKey) {
      throw new NoPublicKey();
    }

    return {
      name: this.csrfHeaderName,
      value: btoa(Md5.hashStr(environment.privateKey + this.publicKey) as string)
    };
  }
}
