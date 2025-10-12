import { Component, OnInit, inject } from '@angular/core';
import { ActivatedRoute, Router, RouterLink } from '@angular/router';
import { ArticleService } from 'src/app/_services/article.service';
import { Article } from 'src/app/_models/article.model';
import { CommonModule } from '@angular/common';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';
import { MatCardModule } from '@angular/material/card';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';

@Component({
  selector: 'app-article-details',
  standalone: true,
  imports: [
    CommonModule,
    RouterLink,
    MatProgressSpinnerModule,
    MatCardModule,
    MatButtonModule,
    MatIconModule,
  ],
  templateUrl: './article-details.component.html',
  styleUrls: ['./article-details.component.scss'],
})
export class ArticleDetailsComponent implements OnInit {
  private route = inject(ActivatedRoute);
  private articleService = inject(ArticleService);
  private router = inject(Router);

  article: Article | undefined;
  isLoading = true;

  ngOnInit() {
    const articleId = this.route.snapshot.paramMap.get('id');
    if (articleId) {
      this.articleService.getArticleById(articleId).subscribe((data) => {
        const article = data as Article;
        if (article.source?.name === 'Abahlali baseMjondolo') {
          const words = article.contentBody.split(' ');
          article.contentBody = words.slice(7).join(' ');
        }

        const continueReadingRegex = /continue reading/gi;
        const match = continueReadingRegex.exec(article.contentBody);

        if (match) {
          const index = match.index;
          const textBefore = article.contentBody.substring(0, index).trim();
          if (!textBefore.endsWith('<br>') && !textBefore.endsWith('</p>')) {
            article.contentBody = article.contentBody.replace(
              match[0],
              `<br><br>${match[0]}`
            );
          }
        }

        this.article = article;
        this.isLoading = false;
      });
    }
  }

  goToEdit() {
    this.router.navigate(['/articles', this.article!.id, 'edit']);
  }
}