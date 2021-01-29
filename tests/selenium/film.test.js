const { createBrowser, base_url } = require('./browser');
const { By, until } = require('selenium-webdriver');

let browser;
beforeAll(async () => browser = await createBrowser());
afterAll(async () => await browser.quit());

describe('Film Page', () => {
  test('it should have a film id', async () => {
    await browser.get(base_url + "/#/film/t000000001");
    browser.wait(until.elementLocated(By.id('film-id'), 100000));
    const id = await browser.findElement(By.id('film-id')).getText()
    expect(id.trim()).toBe('id: t000000001');
  });
  test('it should have a title', async () => {
    await browser.get(base_url + "/#/film/t000000001");
    const title = await browser.findElement(By.id('film-title')).getText()
    expect(title.trim()).toBe('title: Sample Film');
  });
});
