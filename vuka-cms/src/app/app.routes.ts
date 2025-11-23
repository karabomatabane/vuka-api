import { Routes } from '@angular/router';
import { ArticlesComponent } from './pages/articles/articles.component';
import { SourcesComponent } from './pages/sources/sources.component';
import { DirectoriesComponent } from './pages/directories/directories.component';
import { RolesAndPermissionsComponent } from './pages/roles-and-permissions/roles-and-permissions.component';
import { ArticleDetailsComponent } from './pages/article-details/article-details.component';
import { ArticleEditComponent } from './pages/article-edit/article-edit.component';
import { SourceEditComponent } from './pages/source-edit/source-edit.component';
import { LoginComponent } from './pages/login/login.component';
import { authGuard } from './_helpers/auth.guard';
import { DirectoryCategoryComponent } from './pages/directories/directory-category/directory-category.component';
import { NewsletterComponent } from './pages/newsletter/newsletter.component';

export const routes: Routes = [
  {
    path: '',
    pathMatch: 'full',
    redirectTo: 'articles',
  },
  {
    path: 'login',
    component: LoginComponent,
  },
  {
    path: 'articles',
    component: ArticlesComponent,
    canActivate: [authGuard],
  },
  {
    path: 'articles/:id',
    component: ArticleDetailsComponent,
    canActivate: [authGuard],
  },
  {
    path: 'articles/:id/edit',
    component: ArticleEditComponent,
    canActivate: [authGuard],
  },
  {
    path: 'sources',
    component: SourcesComponent,
    canActivate: [authGuard],
  },
  {
    path: 'sources/new',
    component: SourceEditComponent,
    canActivate: [authGuard],
  },
  {
    path: 'sources/:id/edit',
    component: SourceEditComponent,
    canActivate: [authGuard],
  },
  {
    path: 'directories',
    component: DirectoriesComponent,
    canActivate: [authGuard],
  },
  {
    path: 'directories/:id',
    component: DirectoryCategoryComponent,
    canActivate: [authGuard],
  },
  {
    path: 'roles-and-permissions',
    component: RolesAndPermissionsComponent,
    canActivate: [authGuard],
  },
  {
    path: 'newsletter',
    component: NewsletterComponent,
    canActivate: [authGuard],
  }
];
