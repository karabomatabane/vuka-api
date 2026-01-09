import { BaseModel } from './base.model';

export interface NewsletterSubscriber extends BaseModel {
  preferredName: string;
  email: string;
  phoneNumber: string;
}
