import { CsrfService } from './csrf';
import {Md5} from 'ts-md5';
import {environment} from '../../../../environments/environment';
import {InvalidCSRFHeader, NoPublicKey} from './types';

describe('CsrfService', () => {
  let service: CsrfService;

  beforeEach(() => {
    service = new CsrfService();
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });

  it('generating csrf header correctly', () => {
    const testPublicKey = 'test';
    service.setPublicKey(btoa(testPublicKey));

    const csrfHeader = service.getCSRFHeader();

    expect(csrfHeader.name).toBe(service.csrfHeaderName);
    expect(csrfHeader.value).toBe(btoa(Md5.hashStr(environment.privateKey + testPublicKey) as string));
  });

  it('set invalid public key', () => {
    expect(() => service.setPublicKey('hello')).toThrow(new InvalidCSRFHeader());
  });

  it('try to get csrf header with bad service setup', () => {
    expect(() => service.getCSRFHeader()).toThrow(new NoPublicKey());
  });
});
