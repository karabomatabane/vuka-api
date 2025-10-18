import { BaseModel } from './base.model';

export interface ArticleImage extends BaseModel {
  articleId: string;
  isMain: boolean;
  url: string;
  altText?: string;
}
