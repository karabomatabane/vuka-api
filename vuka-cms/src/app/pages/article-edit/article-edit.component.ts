import { Component, OnInit, inject } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormBuilder, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatButtonModule } from '@angular/material/button';
import { MatSlideToggleModule } from '@angular/material/slide-toggle';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';
import { ArticleService } from 'src/app/_services/article.service';
import { Article } from 'src/app/_models/article.model';

@Component({
  selector: 'app-article-edit',
  standalone: true,
  imports: [
    CommonModule,
    ReactiveFormsModule,
    MatFormFieldModule,
    MatInputModule,
    MatButtonModule,
    MatSlideToggleModule,
    MatProgressSpinnerModule,
  ],
  templateUrl: './article-edit.component.html',
  styleUrls: ['./article-edit.component.scss'],
})
export class ArticleEditComponent implements OnInit {
  private route = inject(ActivatedRoute);
  private router = inject(Router);
  private fb = inject(FormBuilder);
  private articleService = inject(ArticleService);

  article: Article | undefined;
  editForm!: FormGroup;
  isLoading = true;
  articleId!: string;

  ngOnInit() {
    this.articleId = this.route.snapshot.paramMap.get('id')!;
    this.editForm = this.fb.group({
      title: ['', Validators.required],
      contentBody: ['', Validators.required],
      isFeatured: [false],
    });

    if (this.articleId) {
      this.articleService.getArticleById(this.articleId).subscribe((data) => {
        this.article = data as Article;
        this.editForm.patchValue(this.article);
        this.isLoading = false;
      });
    }
  }

  save() {
    if (this.editForm.valid) {
      this.isLoading = true;
      const updatedArticle = { ...this.article, ...this.editForm.value };
      this.articleService.updateArticle(this.articleId, updatedArticle).subscribe(() => {
        this.router.navigate(['/articles', this.articleId]);
      });
    }
  }

  cancel() {
    this.router.navigate(['/articles', this.articleId]);
  }
}