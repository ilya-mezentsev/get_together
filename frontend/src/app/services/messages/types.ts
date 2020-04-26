import {ErrorServerResponse} from '../../types/models';

export interface Message {
  chatId: number;
  senderId: number;
  text: string;
  sendingTime?: Date;
}

export interface IMessagesObserver {
  onNewMessage(message: Message): void;
  onError(error: ErrorServerResponse): void;
}

export interface IMessagesObservable {
  connect(observer: IMessagesObserver): Promise<void>;
  send(message: Message): void;
}

export class ConnectionIsNotEstablished extends Error {}
