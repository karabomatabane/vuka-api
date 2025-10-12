import { BaseModel } from './base.model';

export interface Source extends BaseModel {
  name: string;
  websiteUrl: string;
  rssFeedUrl: string;
}
