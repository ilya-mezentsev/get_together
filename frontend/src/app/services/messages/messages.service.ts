import { Injectable } from '@angular/core';
import {IMessagesObservable, IMessagesObserver, Message, ConnectionIsNotEstablished} from './types';
import {environment} from '../../../environments/environment';
import {ErrorServerResponse} from '../../types/models';

@Injectable({
  providedIn: 'root'
})
export class MessagesService implements IMessagesObservable {
  private connectionEstablished = false;
  private readonly messagesObservers: Set<IMessagesObserver> = new Set<IMessagesObserver>();
  private readonly ws = new WebSocket(`ws://${environment.domain}/api/ws/`);

  constructor() {
    this.ws.onopen = () => this.connectionEstablished = true;
    this.ws.onerror = (error: Event) => this.onError(error as unknown as ErrorServerResponse);
    this.ws.onclose = () => this.connectionEstablished = false;
    this.ws.onmessage = (messageEvent: MessageEvent) => this.onMessage(messageEvent.data);
  }

  connect(observer: IMessagesObserver): Promise<void> {
    this.messagesObservers.add(observer);

    if (this.connectionEstablished) {
      return Promise.resolve();
    }

    return new Promise<void>(resolve => {
      this.ws.onopen = () => {
        this.connectionEstablished = true;
        return resolve();
      };
    });
  }

  send(message: Message): void {
    if (!this.connectionEstablished) {
      throw new ConnectionIsNotEstablished();
    }

    this.ws.send(MessagesService.stringify(message));
  }

  private onError(error: ErrorServerResponse): void {
    for (const observer of this.messagesObservers) {
      observer.onError(error);
    }
  }

  private onMessage(message: Message): void {
    for (const observer of this.messagesObservers) {
      observer.onNewMessage(message);
    }
  }

  private static stringify(message: Message): string {
    return JSON.stringify({
      chat_id: message.chatId,
      sender_id: message.senderId,
      text: message.text
    });
  }
}
