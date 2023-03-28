const WORDS_PER_MIN = 275;
const GLYPHS_PER_MIN = 500;

const IMAGE_READ_TIME = 12; // in seconds
const IMAGE_TAGS = ['img', 'Image'];

// Chinese / Japanese / Korean
const GLYPHS_PATTERN = '[\u3040-\u30ff\u3400-\u4dbf\u4e00-\u9fff\uf900-\ufaff\uff66-\uff9f]';

export function wordsCount(text: string) {
    const pattern = '\\w+';
    const reg = new RegExp(pattern, 'g');
    return (text.match(reg) || []).length;
}

export function glyphs(text: string) {
    const reg = new RegExp(GLYPHS_PATTERN, 'g');
    const count = (text.match(reg) || []).length;
    const remainingText = text.replace(reg, '');
    return {
        count,
        remainingText,
    };
}

export function imageCount(imageTags: string[], text: string) {
    const combinedImageTags = imageTags.join('|');
    const pattern = `<(${combinedImageTags})([\\w\\W]+?)[\\/]?>`;
    const reg = new RegExp(pattern, 'g');
    return (text.match(reg) || []).length;
}

export function imageReadTime(imageReadTime = IMAGE_READ_TIME, tags = IMAGE_TAGS, string) {
    let seconds = 0;
    const count = imageCount(tags, string);

    if (count > 10) {
        seconds = ((count / 2) * (imageReadTime + 3)) + (count - 10) * 3; // n/2(a+b) + 3 sec/image
    } else {
        seconds = (count / 2) * (2 * imageReadTime + (1 - count)); // n/2[2a+(n-1)d]
    }
    return {
        time: seconds / 60,
        count,
    };
}

export function stripTags(text: string) {
    const pattern = '<\\w+(\\s+("[^"]*"|\\\'[^\\\']*\'|[^>])+)?>|<\\/\\w+>';
    const reg = new RegExp(pattern, 'gi');
    return text.replace(reg, '');
}

export function stripWhitespace(text: string) {
    return text.replace(/^\s+/, '').replace(/\s+$/, '');
}

export function humanizeTime(time: number) {
    if (time < 0.5) {
        return 'less than a minute';
    } if (time >= 0.5 && time < 1.5) {
        return '1 minute';
    }
    return `${Math.ceil(time)} minutes`;
}

export interface ReadTimeOptions {
    wordsPerMin: number;
    glyphsPerMin: number;
    imageReadTime: number;
    imageTags: string[];
}

export function processText(text: string, options: Partial<ReadTimeOptions> = {}) {
    options = {
        wordsPerMin: WORDS_PER_MIN,
        glyphsPerMin: GLYPHS_PER_MIN,
        imageReadTime: IMAGE_READ_TIME,
        imageTags: IMAGE_TAGS,
        ...options
    };
    text = stripTags(text);
    text = stripWhitespace(text);

    const {count: glyphCount, remainingText} = glyphs(text);
    const wordCount = wordsCount(remainingText);
    const glyphTime = glyphCount / options.wordsPerMin;
    const wordTime = wordCount / options.wordsPerMin;

    const humanizedTime = humanizeTime(glyphTime + wordTime);
    return {
        glyphCount,
        wordCount,
        glyphTime,
        wordTime,
        humanizedTime,
    };
}
