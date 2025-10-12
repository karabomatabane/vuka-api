import { Article } from './article.model';
import { BaseModel } from './base.model';

export interface Region extends BaseModel {
  name: string;
  slug: string;
  articles: Article[];
}
