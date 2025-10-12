import { Component, computed, signal, effect } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { MatToolbarModule } from '@angular/material/toolbar';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { MatSidenavModule } from '@angular/material/sidenav';
import { XSidenavComponent } from './components/x-sidenav/x-sidenav.component';
import { Title } from '@angular/platform-browser';
import { MatTooltipModule } from '@angular/material/tooltip';

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [
    RouterOutlet,
    MatToolbarModule,
    MatButtonModule,
    MatIconModule,
    MatSidenavModule,
    XSidenavComponent,
    MatTooltipModule,
  ],
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss'],
})
export class AppComponent {
  title = 'Vuka CMS';
  collapsed = signal(false);
  isDarkTheme = signal(true);

  constructor(private titleService: Title) {
    this.titleService.setTitle(this.title);
    document.body.classList.add('dark-theme');

    effect(() => {
      if (this.collapsed()) {
        document.body.classList.add('sidenav-collapsed');
      } else {
        document.body.classList.remove('sidenav-collapsed');
      }
    });
  }

  sidebarWidth = computed(() => (this.collapsed() ? '65px' : '250px'));

  toggleTheme() {
    this.isDarkTheme.set(!this.isDarkTheme());
    if (this.isDarkTheme()) {
      document.body.classList.add('dark-theme');
    } else {
      document.body.classList.remove('dark-theme');
    }
  }
}
