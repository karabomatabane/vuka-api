import { BaseModel } from './base.model';
import { Category } from './category.model';
import { Region } from './region.model';
import { Source } from './source.model';

export interface Article extends BaseModel {
  title: string;
  originalUrl: string;
  contentBody: string;
  publishedAt: string;
  isFeatured: boolean;
  sourceId: string;
  source: Source;
  regionID: string | null;
  region: Region;
  categories: Category[] | null;
}
