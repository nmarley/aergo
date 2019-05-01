/**
 * @file    ir_sgmt.c
 * @copyright defined in aergo/LICENSE.txt
 */

#include "common.h"

#include "util.h"

#include "ir_sgmt.h"

static void
sgmt_extend(ir_sgmt_t *sgmt)
{
    sgmt->cap += SGMT_INIT_CAPACITY;

    sgmt->lens = xrealloc(sgmt->lens, sizeof(BinaryenIndex) * sgmt->cap);
    sgmt->addrs = xrealloc(sgmt->addrs, sizeof(BinaryenIndex) * sgmt->cap);
    sgmt->datas = xrealloc(sgmt->datas, sizeof(char *) * sgmt->cap);
}

static int
sgmt_lookup(ir_sgmt_t *sgmt, void *ptr, uint32_t len)
{
    int i;

    for (i = 0; i < sgmt->size; i++) {
        if (sgmt->lens[i] == len && memcmp(sgmt->datas[i], ptr, len) == 0)
            return sgmt->addrs[i];
    }

    return -1;
}

int
sgmt_add_str(ir_sgmt_t *sgmt, char *str)
{
    int i, j;
    int len;
    char *esc_str;

    ASSERT(str != NULL);

    len = strlen(str);

    if (strchr(str, '\\') == NULL) {
        return sgmt_add_raw(sgmt, str, len + 1);
    }

    esc_str = xmalloc(len + 1);

    for (i = 0, j = 0; i < len; i++) {
        if (str[i] == '\\' && isesc(str[i + 1]))
            esc_str[j++] = etoc(str[i++ + 1]);
        else
            esc_str[j++] = str[i];
    }

    ASSERT2(j < len, j, len);

    esc_str[j++] = '\0';

    return sgmt_add_raw(sgmt, esc_str, j);
}

int
sgmt_add_raw(ir_sgmt_t *sgmt, void *ptr, uint32_t len)
{
    int addr;

    ASSERT(ptr != NULL);
    ASSERT1(len > 0, len);

    addr = sgmt_lookup(sgmt, ptr, len);
    if (addr >= 0)
        return addr;

    if (sgmt->size >= sgmt->cap)
        sgmt_extend(sgmt);

    /* TODO: Apply proper alignment */
    if (len > 4)
        sgmt->offset = ALIGN64(sgmt->offset);
    else
        sgmt->offset = ALIGN32(sgmt->offset);

    addr = sgmt->offset;

    sgmt->lens[sgmt->size] = len;
    sgmt->addrs[sgmt->size] = addr;
    sgmt->datas[sgmt->size] = ptr;

    sgmt->size++;
    sgmt->offset += len;

    return addr;
}

/* end of ir_sgmt.c */