const { createBrowser, base_url } = require('./browser');
const { By, until, Key } = require('selenium-webdriver');
const { expect } = require('@jest/globals');

jest.retryTimes(3);
jest.setTimeout(30 * 1000);

let browser;
beforeAll(async () => browser = await createBrowser());
afterAll(async () => await browser.quit());

describe('Search Results page', () => {
  test('it should display results for films', async () => {
    await browser.get(base_url);
    await browser.wait(until.elementLocated(By.id('global-search-query'), 500));
    await browser.findElement(By.id('global-search-query')).sendKeys("Sample");
    browser.findElement(By.id('global-search-btn')).click();
    await browser.wait(until.elementsLocated(By.className('results-grid')), 500);
    let cardTitle = await browser.findElement(By.className('card-title')).getText();

    expect(cardTitle.trim()).toBe('Sample Film');
  });
  test('it should display results for people', async () => {
    await browser.get(base_url);
    await browser.wait(until.elementLocated(By.id('global-search-query'), 500));
    await browser.findElement(By.id('global-search-query')).sendKeys("John");
    browser.findElement(By.id('global-search-btn')).click();
    await browser.wait(until.elementsLocated(By.className('results-grid')), 500);
    let cardTitle = await browser.findElement(By.className('card-title')).getText();

    expect(cardTitle.trim()).toBe('John Doe');
  });
  test('it should search by enter', async () => {
    await browser.get(base_url);
    await browser.wait(until.elementLocated(By.id('global-search-query'), 500));
    await browser.findElement(By.id('global-search-query')).sendKeys("Sample", Key.ENTER);
    await browser.wait(until.elementsLocated(By.className('results-grid')), 500);
    let cardTitle = await browser.findElement(By.className('card-title')).getText();

    expect(cardTitle.trim()).toBe('Sample Film');
  });
  test('it should do nothing on empty query', async () => {
    await browser.get(base_url);
    await browser.wait(until.elementLocated(By.id('global-search-query'), 500));
    await browser.findElement(By.id('global-search-btn')).click();
    const url = await browser.getCurrentUrl();

    let match_url = base_url;
    if (!match_url.endsWith('/')) match_url += '/';
    
    expect(url).toBe(match_url);
  });
});
