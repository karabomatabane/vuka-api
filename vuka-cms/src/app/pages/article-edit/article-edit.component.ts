import { Component, OnInit, inject } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormBuilder, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';

import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatButtonModule } from '@angular/material/button';
import { MatSlideToggleModule } from '@angular/material/slide-toggle';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';
import { ArticleService } from 'src/app/_services/article.service';
import { Article } from 'src/app/_models/article.model';
import { CategoryService } from 'src/app/_services/category.service';
import { Category } from 'src/app/_models/category.model';
import { MatSelectModule } from '@angular/material/select';

@Component({
  selector: 'app-article-edit',
  standalone: true,
  imports: [
    ReactiveFormsModule,
    MatFormFieldModule,
    MatInputModule,
    MatButtonModule,
    MatSlideToggleModule,
    MatProgressSpinnerModule,
    MatSelectModule
],
  templateUrl: './article-edit.component.html',
  styleUrls: ['./article-edit.component.scss'],
})
export class ArticleEditComponent implements OnInit {
  private route = inject(ActivatedRoute);
  private router = inject(Router);
  private fb = inject(FormBuilder);
  private articleService = inject(ArticleService);
  private categoryService = inject(CategoryService);

  article: Article | undefined;
  editForm!: FormGroup;
  isLoading = true;
  articleId!: string;
  categories: Category[] = [];

  ngOnInit() {
    this.articleId = this.route.snapshot.paramMap.get('id')!;
    this.editForm = this.fb.group({
      title: ['', Validators.required],
      summary: ['', Validators.required],
      isFeatured: [false],
      categoryIds: [[]],
    });

    this.categoryService.getCategories().subscribe((data) => {
      this.categories = data;
    });

    if (this.articleId) {
      this.articleService.getArticleById(this.articleId).subscribe((data) => {
        this.article = data as Article;
        this.editForm.patchValue({
          ...this.article,
          categoryIds: this.article.categories?.map((c) => c.id) || [],
        });
        this.isLoading = false;
      });
    }
  }

  save() {
    if (this.editForm.valid) {
      this.isLoading = true;
      const formValue = this.editForm.value;
      
      // Only send the fields that should be updated
      const updates = {
        title: formValue.title,
        summary: formValue.summary,
        contentBody: formValue.contentBody,
        isFeatured: formValue.isFeatured,
        categoryIds: formValue.categoryIds
      };
      
      this.articleService
        .updateArticle(this.articleId, updates)
        .subscribe(() => {
          this.router.navigate(['/articles', this.articleId]);
        });
    }
  }

  cancel() {
    this.router.navigate(['/articles', this.articleId]);
  }
}