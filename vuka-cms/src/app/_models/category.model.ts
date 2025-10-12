import { Article } from './article.model';
import { BaseModel } from './base.model';

export interface Category extends BaseModel {
  name: string;
  articles: Article[];
}
