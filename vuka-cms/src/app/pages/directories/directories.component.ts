import { Component, inject } from '@angular/core';
import { DirectoryService } from 'src/app/_services/directory.service';
import { MatIcon, MatIconModule } from '@angular/material/icon';
import { MatButton } from '@angular/material/button';
import { MatDialog } from '@angular/material/dialog';
import { DirectoryFormDialogComponent } from 'src/app/components/directory-form-dialog/directory-form-dialog.component';
import { AnimationOptions, LottieComponent } from 'ngx-lottie';
import { AnimationItem } from 'lottie-web';
import { DirectoryOverview } from 'src/app/_models/directory-category.model';
import { MatListModule } from '@angular/material/list';
import { Router } from '@angular/router';


@Component({
  standalone: true,
  selector: 'app-directories',
  imports: [
    MatIcon,
    MatButton,
    LottieComponent,
    MatListModule,
    MatIcon
],
  templateUrl: './directories.component.html',
  styleUrl: './directories.component.scss',
})
export class DirectoriesComponent {
  constructor(
    private directoryService: DirectoryService,
    private router: Router,
  ) {}
  readonly dialog = inject(MatDialog);
  overview?: DirectoryOverview;

  options: AnimationOptions = {
    path: '/lottie/no-data.json',
  };

  ngOnInit() {
    this.directoryService.getOverview().subscribe((data) => {
      this.overview = data;
    });
  }

  animationCreated(animationItem: AnimationItem): void {
    console.log(animationItem);
  }

  onAddCategory() {
    const dialogRef = this.dialog.open(DirectoryFormDialogComponent);

    dialogRef.afterClosed().subscribe((result) => {
      if (result) {
        // Broadcast event to Navbar to refresh directories
      }
    });
  }

  openDirectory(id: string) {
    this.router.navigate([`/directories/${id}`]);
  }
}
