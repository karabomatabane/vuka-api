import { Routes } from '@angular/router';
import { DashboardComponent } from './pages/dashboard/dashboard.component';
import { ArticlesComponent } from './pages/articles/articles.component';
import { SourcesComponent } from './pages/sources/sources.component';
import { DirectoriesComponent } from './pages/directories/directories.component';
import { RolesAndPermissionsComponent } from './pages/roles-and-permissions/roles-and-permissions.component';
import { ArticleDetailsComponent } from './pages/article-details/article-details.component';
import { ArticleEditComponent } from './pages/article-edit/article-edit.component';

export const routes: Routes = [
  {
    path: '',
    pathMatch: 'full',
    redirectTo: 'dashboard',
  },
  {
    path: 'dashboard',
    component: DashboardComponent,
  },
  {
    path: 'articles',
    component: ArticlesComponent,
  },
  {
    path: 'articles/:id',
    component: ArticleDetailsComponent,
  },
  {
    path: 'articles/:id/edit',
    component: ArticleEditComponent,
  },
  {
    path: 'sources',
    component: SourcesComponent,
  },
  {
    path: 'directories',
    component: DirectoriesComponent,
  },
  {
    path: 'roles-and-permissions',
    component: RolesAndPermissionsComponent,
  },
];
