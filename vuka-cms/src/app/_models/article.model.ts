export interface Article {
  id: string;
  createdAt: string;
  updatedAt: string;
  deletedAt: string | null;
  title: string;
  originalUrl: string;
  contentBody: string;
  publishedAt: string;
  isFeatured: boolean;
  sourceId: string;
  source: {
    id: string;
    name: string;
    websiteUrl: string;
    rssFeedUrl: string;
  };
  regionID: string | null;
  region: {
    id: string;
    name: string;
    slug: string;
  };
  Categories: any[] | null; // Replace 'any' with a proper Category interface if available
}